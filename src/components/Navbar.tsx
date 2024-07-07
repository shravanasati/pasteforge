import React from "react"
import { ThemeType, themes } from "../themes"
import { AlignLeft, Palette } from "lucide-react";
import logo from "../assets/logo.png"

type NavbarProps = {
  setTheme: (theme: ThemeType) => void;
}

// todo add language dropdown

export const Navbar: React.FC<NavbarProps> = ({ setTheme }) => {
  const [selectedTheme, setSelectedTheme] = React.useState<ThemeType>(themes.dark)
  const [isDropdownOpen, setIsDropdownOpen] = React.useState(false)

  const handleThemeChange = (theme: ThemeType) => {
    setTheme(theme);
    setSelectedTheme(theme);
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
        <a className="btn btn-ghost text-xl hover:underline">pasteforge</a>
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
          <div tabIndex={0} role="button" className="btn btn-ghost" onClick={toggleDropdown}>
            <Palette size={20} />
            {selectedTheme}
          </div>
          {isDropdownOpen && (
            <ul
              tabIndex={0}
              className="menu menu-sm dropdown-content bg-base-100 rounded-box z-[1] mt-3 w-52 p-2 shadow"
            >
              {Object.keys(themes).map((theme) => (
                <li key={theme}>
                  <a onClick={() => {
                    handleThemeChange(theme as ThemeType)
                  }}>{theme}</a>
                </li>
              ))}
            </ul>)}
        </div>
        <a className="btn">save</a>
      </div>
    </div>
  );
};