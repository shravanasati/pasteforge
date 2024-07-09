import { useEffect, useState } from 'react'
import './App.css'
import { Navbar } from "./components/Navbar"
import { Footer } from "./components/Footer"
import { CodeEditor } from './components/CodeEditor'
import { Options } from './components/Options'

function App() {
  const [lang, setLang] = useState("plain")
      
  let ce = <CodeEditor language={lang} />
  useEffect(
    () => {
      ce = <CodeEditor language={lang} />
    },
    [lang]
  )

  return (
    <>
      <Navbar setLang={setLang} />
      {ce}
      <Options />
      <Footer/>
    </>
  )
}

export default App
