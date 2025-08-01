import { Link, Outlet, useNavigate } from "react-router-dom"
import { useState } from "react"
import Alert from "./Alert"

const App = () => {

    const [jwtToken, setJwtToken] = useState("")
    const [alertMessage, setAlertMessage] = useState("")
    const [alertClassName, setAlertClassName] = useState("d-none")

    const navigate = useNavigate()

    const logout = () => {
        setJwtToken("")
        navigate("/login")
    }

    return (
        <>
            <div className="container">
                <div className="row">
                    <div className="col-md-8">
                        <h1>Movies Mania</h1>
                    </div>
                    <div className="col text-end">
                        {
                            jwtToken === ""
                                ? <Link to="/login" className="badge bg-success">Login</Link>
                                : <Link to="#!" className="badge bg-danger" onClick={logout}>Logout</Link>
                        }
                    </div>
                </div>
                <hr />
                <div className="row">
                    <div className="col-md-2">
                        <ul className="list-group ">
                            <Link to="/home" className="list-group-item shadow-sm p-3 mb-2 rounded fw-bold text-dark list-group-item-action">Home</Link>
                            <Link to="/movies" className="list-group-item shadow-sm p-3 mb-2 rounded fw-bold text-dark list-group-item-action">Movies</Link>
                            <Link to="/explore" className="list-group-item shadow-sm p-3 mb-2 rounded fw-bold text-dark list-group-item-action">Explore</Link>
                            <Link to="/genres" className="list-group-item shadow-sm p-3 mb-2 rounded fw-bold text-dark list-group-item-action">Genres</Link>
                            <Link to="/exit" className="list-group-item shadow-sm p-3 mb-2 rounded fw-bold text-dark list-group-item-action">Exit</Link>
                        </ul>
                    </div>
                    <div className="col-md-10">
                        <Alert message={alertMessage} className={alertClassName} />
                        <Outlet context={{
                            jwtToken, setJwtToken, setAlertClassName, setAlertMessage
                        }} />
                    </div>
                </div>
            </div>
        </>
    )
}

export default App