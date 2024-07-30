import { useEffect, useState } from 'react'
import './App.css'
import { Navbar } from "./components/Navbar"
import { Footer } from "./components/Footer"
import { CodeEditor } from './components/CodeEditor'
import { PasteSettings } from './components/PasteSettings'

function App() {
  const [lang, setLang] = useState("plain")
  const [code, setCode] = useState("")

  let ce = <CodeEditor language={lang} code={code} setCode={setCode}/>
  useEffect(
    () => {
      ce = <CodeEditor language={lang} code={code} setCode={setCode}/>
    },
    [lang]
  )

  return (
    <>
      <Navbar setLang={setLang} />
      {ce}
      <PasteSettings language={lang} code={code} />
      <Footer />
    </>
  )
}

export default App
