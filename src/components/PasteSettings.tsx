import { CircleX } from "lucide-react"
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

export function PasteSettings() {
  const [error, setError] = useState("")

  const expirationNumRef = useRef<HTMLInputElement>(null)
  const expirationDurationRef = useRef<HTMLSelectElement>(null)
  const visibilityRef = useRef<HTMLSelectElement>(null)

  function expirationDropdown() {
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
            } else {
              setError("")
            }
          }} />
          <select className="select select-secondary w-full max-w-xs" ref={expirationDurationRef} >
            <option disabled selected>duration</option>
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

  function visibility() {
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

  function settingsValidationError() {
    return <div role="alert" className="alert rounded alert-error w-auto">
      <CircleX size={20} />
      <span>{error}</span>
    </div>
  }

  const [loading, setLoading] = useState(false)

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    if (error) {
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

    // todo send to server
    setLoading(true)
    console.log("save button clicked")
    setTimeout(() => setLoading(false), 2000)
  }

  useEffect(() => {
    const settingsJSON = localStorage.getItem("settings")
    let settings: Settings;

    if (settingsJSON) {
      settings = JSON.parse(settingsJSON) as Settings
    } else {
      settings = defaultSettings
    }

    expirationNumRef.current!.value = settings.expiration.toString()
    expirationDurationRef.current!.value = settings.duration
    visibilityRef.current!.value = settings.visibility
  }, [])

  return <form onSubmit={handleSubmit}>
    <h1 className="m-2 p-2">paste settings</h1>
    <div className="flex flex-row justify-start flex-wrap items-center ml-4">
      {expirationDropdown()}
      {visibility()}
      {error == "" ? null : settingsValidationError()}
    </div>
    {!loading ? <input type="submit" value="save" className="btn m-4" /> : <span className="btn m-4 loading loading-dots text-secondary" />}
  </form>
}