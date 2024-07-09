import { CircleX } from "lucide-react"
import { useState } from "react"


export function Options() {
  const [error, setError] = useState("")

  function expirationDropdown() {
    return (
      <div className="flex items-center">
        <div tabIndex={0} role="button" className="text-primary m-2">expiration: </div>
        <div
          tabIndex={0}
          className="text-primary-content z-[1] w-64 p-2 flex flex-row m-2">
          <input type="number" min={0} placeholder="x" className="input input-bordered w-full max-w-xs" onChange={(e) => {
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
          <select className="select select-secondary w-full max-w-xs">
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
        <select className="select select-secondary w-full max-w-xs">
          <option selected>public</option>
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

  return <>
    <h1 className="m-2 p-2">paste settings</h1>
    <div className="flex flex-row justify-start flex-wrap items-center space-x-4">
      {expirationDropdown()}
      {visibility()}
      {error == "" ? null : settingsValidationError()}
    </div>
  </>
}