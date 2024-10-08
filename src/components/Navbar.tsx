import React from "react"
import { Code, AlignLeft } from "lucide-react";
import logo from "../assets/logo.png"
import { languages } from "../languages";

type NavbarProps = {
  setLang: (lang: string) => void;
}

function MiddleNavbarItems() {
  return (
    <>
      <li>
        <a>API</a>
      </li>
      <li>
        <a>CLI</a>
      </li>
      <li>
        <a>VSCode</a>
      </li>
    </>
  )
}


export const Navbar: React.FC<NavbarProps> = ({ setLang }) => {
  const [selectedLang, setSelectedLang] = React.useState("plain")
  const [isDropdownOpen, setIsDropdownOpen] = React.useState(false)

  const handleLangChange = (lang: string) => {
    setLang(lang);
    setSelectedLang(lang);
    setIsDropdownOpen(false);
  }
  const toggleDropdown = () => setIsDropdownOpen(!isDropdownOpen)

  return (
    <div className="navbar bg-base-100">
      <div className="navbar-start">
        <div className="dropdown">
          <div tabIndex={0} role="button" className="btn btn-ghost lg:hidden">
            <AlignLeft size={24} color="black" />
          </div>
          <ul
            tabIndex={0}
            className="menu menu-sm dropdown-content bg-base-100 rounded-box z-[1] mt-3 w-52 p-2 shadow"
          >
          {MiddleNavbarItems()}
          </ul>
        </div>
        <a className="btn btn-ghost text-xl hover:underline text-black rounded">
          <img src={logo} alt="logo" height={32} width={32} /> pasteforge
        </a>
      </div>
      <div className="navbar-center hidden lg:flex">
        <ul className="menu menu-horizontal px-1">
          {MiddleNavbarItems()}
        </ul>
      </div>
      <div className="navbar-end">
        <div className="dropdown">
          <div tabIndex={0} role="button" className="btn btn-ghost text-black" onClick={toggleDropdown}>
            <Code size={20} />
            {selectedLang}
          </div>
          {isDropdownOpen && (
            <ul
              tabIndex={0}
              className="menu menu-sm dropdown-content bg-base-100 rounded-box z-[1] mt-3 w-28 p-2 shadow"
            >
              {Object.keys(languages).map((lang) => (
                <li key={lang}>
                  <a onClick={() => {
                    handleLangChange(lang)
                  }}>{lang}</a>
                </li>
              ))}
            </ul>)}
        </div>
        <a className="btn mx-1">login</a>
      </div>
    </div>
  );
};