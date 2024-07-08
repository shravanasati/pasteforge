import logo from '../assets/logo.png'

export function Footer() {
  return (
    <footer className="footer bg-neutral text-neutral-content items-center p-4">
      <aside className="grid-flow-col items-center">
        <img src={logo} alt="logo" className="rotate-180" height={38} width={38} />
        <p >Copyright &copy; {new Date().getFullYear()} - all rights reserved by <a className="text-primary" href="https://shravanasati.me">shravan asati</a></p>
      </aside>
      <nav className="grid-flow-col gap-4 md:place-self-center md:justify-self-end">
        <strong>
          social labels are bloat
        </strong>
      </nav>
    </footer>
  )
}
