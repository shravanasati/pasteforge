import React from "react"
import { AlignLeft, Code } from "lucide-react";
import logo from "../assets/logo.png"
import { languages } from "../languages";

type NavbarProps = {
  setLang: (lang: string) => void;
}


export const Navbar: React.FC<NavbarProps> = ({ setLang  }) => {
  const [selectedLang, setSelectedLang] = React.useState(languages.plain)
  const [isDropdownOpen, setIsDropdownOpen] = React.useState(false)

  const handleThemeChange = (lang: string) => {
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
            <AlignLeft size={24} />
          </div>
          <ul
            tabIndex={0}
            className="menu menu-sm dropdown-content bg-base-100 rounded-box z-[1] mt-3 w-52 p-2 shadow"
          >
            <li>
              <a>Item 1</a>
            </li>
            <li>
              <a>Item 3</a>
            </li>
          </ul>
        </div>
        <img src={logo} alt="logo" height={32} width={32} />
        <a className="btn btn-ghost text-xl hover:underline text-black">pasteforge</a>
      </div>
      <div className="navbar-center hidden lg:flex">
        <ul className="menu menu-horizontal px-1">
          <li>
            <a>Item 1</a>
          </li>
          <li>
            <a>Item 3</a>
          </li>
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
              className="menu menu-sm dropdown-content bg-base-100 rounded-box z-[1] mt-3 w-52 p-2 shadow overflow-x-auto h-screen"
            >
              {Object.keys(languages).map((lang) => (
                <li key={lang}>
                  <a onClick={() => {
                    handleThemeChange(lang)
                  }}>{lang}</a>
                </li>
              ))}
            </ul>)}
        </div>
        <a className="btn mx-1">save</a>
        <a className="btn mx-1">login</a>
      </div>
    </div>
  );
};