import { Link, useRouteError } from "react-router-dom";

function GoBack() {
  return (
    <button>
      <Link to={'/'} > Go back to the home page </Link>
    </button>
  );
}

export default function ErrorPage() {
  const error = useRouteError() as { statusText: string, message: string };
  console.error(error);

  const is404 = error.statusText === "Not Found";

  return (
    <div className="h-screen w-screen flex flex-col justify-center items-center">
      <h1 className="p-2 m-4">Oops!</h1>
      <p className="p-2 m-2">{is404 ? "This page doesn't exist." : "Sorry, an unexpected error has occurred."}</p>
      {
        is404 ? <GoBack /> : <p>
          <i>{error.statusText || error.message}</i>
        </p>
      }

    </div>
  );
}