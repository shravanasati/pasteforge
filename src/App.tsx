import { useState } from 'react'
import './App.css'
import { ThemeType, themes } from './themes'
import {Navbar} from "./components/Navbar"


function App() {
  const [theme, setTheme] = useState<ThemeType>(themes.dark)

  return (
    <div data-theme={theme}>
      <Navbar setTheme={setTheme} />
    </div>
  )
}

export default App
