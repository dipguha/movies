import { useState } from "react";
import Input from "./form/input";
import { Navigate, useNavigate, useOutletContext } from "react-router-dom";

const Login = () => {

    //useState returns an array and [] square brackets are used for destructure to get the state variable and the corresponding updatefunction
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    //Custom hook returns objects and curly braces are used to extract specific values or functions
    const {setJwtToken} = useOutletContext();
    const {setAlertMessage} = useOutletContext();
    const {setAlertClassName} = useOutletContext();

    const navigate = useNavigate();
    
    console.log("***** Login.js *****")

    const handleSubmit = (event) => {
        event.preventDefault();
        
        console.log("***** Login.js, email/password *****", email, password)

        if (email === "a@a.com") {
            setJwtToken("abc")
            setAlertClassName("d-non")
            setAlertMessage("")
            navigate("/")
        } else {
            setAlertClassName("alert-danger")
            setAlertMessage("Wrong Credentials, Please try again")
        }
    }

    return (
        <div className="col-md-6 offset-md-3">
            <h2>Login</h2>
            <hr />

            <form onSubmit={handleSubmit}>
                <Input 
                    title="Email Address"
                    type="email"
                    className="form-control"
                    name="email"
                    autoComplete="email-new"
                    onChange={ (event) => setEmail(event.target.value) }
                />

                <Input     
                    title="Password"
                    type="password"
                    className="form-control"
                    name="password"
                    autoComplete="password-new"
                    onChange={ (event) => setPassword(event.target.value) }
                />
                <hr />
                <input type="submit" className="btn btn-primary" value="Login"/>
            </form>
        </div>
    )
}

export default Login;
