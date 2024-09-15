import { CircleX, KeyRoundIcon } from "lucide-react"
import { useEffect, useRef, useState } from "react"

type durationEnum = "minutes" | "hours" | "days" | "months" | "years" | "never"
type visibilityEnum = "public" | "unlisted" | "private"

type Settings = {
  expiration: number
  duration: durationEnum
  visibility: visibilityEnum
}

const defaultSettings: Settings = {
  expiration: 0,
  duration: "never",
  visibility: "public"
}

function validateSettings(settings: any): Settings {
  if (!settings.expiration || (settings.expiration && typeof settings.expiration != "number")) {
    settings.expiration = 0
  }
  if (!settings.duration || typeof settings.duration != "string" || !["minutes", "hours", "days", "months", "years", "never"].includes(settings.duration)) {
    settings.duration = "never"
  }
  if (!settings.visibility || typeof settings.visibility != "string" || !["public", "unlisted", "private"].includes(settings.visibility)) {
    settings.visibility = "public"
  }

  return settings
}

type ExpirationDropdownProps = {
  expirationNumRef: React.RefObject<HTMLInputElement>
  expirationDurationRef: React.RefObject<HTMLSelectElement>
  setError: React.Dispatch<React.SetStateAction<string>>
}

function ExpirationDropdown({ expirationNumRef, expirationDurationRef, setError }: ExpirationDropdownProps) {
  return (
    <div className="flex items-center">
      <div tabIndex={0} role="button" className="text-primary m-2">expiration: </div>
      <div
        tabIndex={0}
        className="text-primary-content z-[1] w-64 p-2 flex flex-row m-2">
        <input type="number" min={0} placeholder="x" ref={expirationNumRef} className="input input-bordered w-full max-w-xs" onChange={(e) => {
          if (!e.target.value) {
            setError("")
            return
          }
          const num = parseInt(e.target.value)
          if (isNaN(num)) {
            setError("expiration time must be a number")
            return
          }
          if (num < 0) {
            setError("expiration time cannot be negative")
          } else if (num == 0 && expirationDurationRef.current!.value != "never") {
            setError("expiration time must be greater than 0")
          } else {
            setError("")
          }
        }} />
        <select className="select select-secondary w-full max-w-xs" ref={expirationDurationRef} onChange={(e) => {
          if (e.target.value != "never" && expirationNumRef.current!.value == "0") {
            setError("expiration time must be greater than 0")
          } else {
            setError("")
          }
        }} >
          <option disabled >duration</option>
          <option>minutes</option>
          <option>hours</option>
          <option>days</option>
          <option>months</option>
          <option>years</option>
          <option selected>never</option>
        </select>
      </div>
    </div>
  )
}

type VisibilityProps = {
  visibilityRef: React.RefObject<HTMLSelectElement>
}

function Visibility({ visibilityRef }: VisibilityProps) {
  return <div className="flex items-center">
    <div tabIndex={0} role="button" className="text-primary m-2">visibility: </div>
    <div
      tabIndex={0}
      className="text-primary-content z-[1] w-64 p-2 flex flex-row m-2">
      <select className="select select-secondary w-full max-w-xs" ref={visibilityRef}>
        <option selected>public</option>
        <option>unlisted</option>
        <option>private</option>
      </select>
    </div>
  </div>
}

function PasswordField({passwordRef}: {passwordRef: React.RefObject<HTMLInputElement>}) {
  return <div className="flex items-center">
    <KeyRoundIcon size={24} className="mx-2 text-primary" />
    <input type="password" placeholder="password" className="input input-bordered input-secondary w-full max-w-xs grow text-black" ref={passwordRef} />
  </div>
}

function SettingsValidationError({ error }: { error: string }) {
  return <div role="alert" className="alert rounded alert-error w-auto m-2">
    <CircleX size={20} />
    <span>{error}</span>
  </div>
}

async function handleSubmit(e: React.FormEvent, error: string, setError: React.Dispatch<React.SetStateAction<string>>, expirationNumRef: React.RefObject<HTMLInputElement>, expirationDurationRef: React.RefObject<HTMLSelectElement>, visibilityRef: React.RefObject<HTMLSelectElement>, setLoading: React.Dispatch<React.SetStateAction<boolean>>, code: string, language: string, passwordRef: React.RefObject<HTMLInputElement>) {
  e.preventDefault()
  if (error) {
    return
  }
  if (!expirationNumRef.current!.value && expirationDurationRef.current!.value != "never") {
    setError("enter a number for expiration time")
    return
  } else {
    setError("")
  }

  // if (!code) {
  //   setError("enter some paste content")
  //   return
  // }

  const expirationNumber = parseInt(expirationNumRef.current!.value, 10)
  if (!expirationNumber) {
    setError("expiration number is not a number")
    return
  }

  const settings: Settings = {
    expiration: parseInt(expirationNumRef.current!.value),
    duration: expirationDurationRef.current!.value as durationEnum,
    visibility: visibilityRef.current!.value as visibilityEnum
  }
  try {
    localStorage.setItem("settings", JSON.stringify(settings))
  } catch (e) {
    console.error("local storage quota exceeded")
  }

  setLoading(true)
  console.log("save button clicked")
  try {
    const reqBody = {
      "content": code,
      "settings": {
        "expiration_number": expirationNumber,
        "expiration_duration": expirationDurationRef.current!.value,
        "visibility": visibilityRef.current!.value,
        "language": language,
        "password": passwordRef.current!.value
      }
    } 
    console.log(reqBody)

    const resp = await fetch("/api/v1/pastes/new", {
      method: "POST",
      body: JSON.stringify(reqBody)
    })
    if (resp.status != 200) {
      console.error("request failed", resp.json())
      setLoading(false)
      return
    }
    console.log("saved")
  } catch (e) {
    console.error("request failed", e)
    setLoading(false)
  }
  setLoading(false)
}

type PasteSettingsProps = {
  language: string
  code: string
}

export function PasteSettings({ language, code }: PasteSettingsProps) {
  const [error, setError] = useState("")
  const [loading, setLoading] = useState(false)

  const expirationNumRef = useRef<HTMLInputElement>(null)
  const expirationDurationRef = useRef<HTMLSelectElement>(null)
  const visibilityRef = useRef<HTMLSelectElement>(null)
  const passwordRef = useRef<HTMLInputElement>(null)

  useEffect(() => {
    const settingsJSON = localStorage.getItem("settings")
    let settings: Settings;

    if (settingsJSON) {
      settings = validateSettings(JSON.parse(settingsJSON))
    } else {
      settings = defaultSettings
    }

    expirationNumRef.current!.value = settings.expiration.toString();
    expirationDurationRef.current!.value = settings.duration
    visibilityRef.current!.value = settings.visibility
  }, [])

  return <form onSubmit={(e) => handleSubmit(e, error, setError, expirationNumRef, expirationDurationRef, visibilityRef, setLoading, code, language, passwordRef)}>
    <h1 className="m-2 p-2">paste settings</h1>
    <div className="flex flex-row justify-start flex-wrap items-center ml-4">
      <ExpirationDropdown expirationNumRef={expirationNumRef} expirationDurationRef={expirationDurationRef} setError={setError} />
      <Visibility visibilityRef={visibilityRef} />
      <PasswordField passwordRef={passwordRef} />
      {error === "" ? null : <SettingsValidationError error={error} />}
    </div>
    {!loading ? <input type="submit" value="save" className="btn m-4" /> : <span className="btn m-4 loading loading-dots text-primary" />}
  </form>
}