import React, {Fragment, useState} from "react";

function HelloWorld(props) {
    const [isTrue, setIsTrue] = useState(false)

    const toggler = () => {
        if (isTrue) {
            setIsTrue(false)
            return 
        }
        setIsTrue(true)
    }
    return (
        <>
            <hr />
            <h1>{props.msg}</h1>
            {
                isTrue ?
                <>
                    <p>I am fucked</p>
                </>
                :
                <>
                    Come on
                </>
            }
            <hr />
            <a href="#!" className="btn btn-outline-secondary" onClick={toggler}>Fuck Me</a>
        </>
    )
}

export default HelloWorld