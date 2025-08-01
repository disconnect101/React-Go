import { useRouteError } from "react-router-dom"

const ErrorElement = () => {
    const error = useRouteError()
    return (
        <>
            <h1>Encountered Error</h1>
            <p>{error.statustext || error.message}</p>
        </>
    )
}

export default ErrorElement