import { useState } from 'react'
import './App.css'
import { Navbar } from "./components/Navbar"
import { languages } from './languages'
import { CodeEditor } from './components/CodeEditor'

function App() {
  const [lang, setLang] = useState(languages.plain)

  return (
    <>
      <Navbar setLang={setLang} />
      <CodeEditor/>
    </>
  )
}

export default App
