import { useEffect, useState } from "react";
import { Link, Outlet, useNavigate } from "react-router-dom";
import Alert from "./components/Alert";

function App() {
  
  console.log("***** App.js *****")
  const [jwtToken, setJwtToken] = useState("");
  const [alertMessage, setAlertMessage] = useState("");
  const [alertClassName, setAlertClassName] = useState("d-none");
  
  const navigate = useNavigate();

  const logout = () => {
    console.log("****** App.js-logout-jwtToken: ",jwtToken)

    const requestOptions = {
      method: "GET",
      credentials: "include"
    }

    //console.log("****** App.js-logout-calling logout: ",jwtToken)
    fetch("/logout", requestOptions)
      .catch( error => {
        console.log("error logout: ", error)
      })
      .finally(() => {
        setJwtToken("")
      })
    
    navigate("/login")
  }

  /* perform side effects (e.g. data fetrching, DOM changes) in functional components, replicate lifecycle methods of class based components like 
  componentDidMount, componentDidUpdate, and componentWillUnmount.
  The effect runs after the initial render and whenever one of the dependencies changes.
  */
  useEffect( () => {
    console.log("****** App.js-useEffect-jwtToken: ", jwtToken)
    if (jwtToken === "") {

      const requestOptions = {
        method: "GET",
        credentials: "include",
      }

      console.log("****** App.js-useEffect-calling /refresh-requestOptions: ", requestOptions)
      fetch(`/refresh`, requestOptions)
        .then( (response) => response.json())
        .then( (data) => {
          if (data.access_token) {
            setJwtToken(data.access_token);
          }
        })
        .catch(error => {
          console.log("user is not logged in: ", error);
        })
    }
  }, [jwtToken])

  return (
    <div className="container">
      <div className="row">
        <div className="col">
          <h1 className="mt-3">Watch a Movie</h1>
        </div>

        <div className="col text-end">
          {jwtToken === ""
            ? <Link to="/Login"><span className="badge bg-success">Login</span></Link>
            : <a href="#!" onClick={logout}><span className="badge bg-danger">Logout</span></a>
          }
        </div>
        <hr className="mb-3"></hr>
      </div>

      <div className="row">
        <div className="col-md-2">
          <nav>
            <div className="list-group">
              <Link to="/" className="list-group-item list-group-item-action">Home</Link>
              <Link to="/movies" className="list-group-item list-group-item-action">Movies</Link>
              <Link to="/genres" className="list-group-item list-group-item-action">Genres</Link>
              {jwtToken !== "" &&
                <>
                  <Link to="/admin/movie/0" className="list-group-item list-group-item-action">Add Movie</Link>
                  <Link to="/manage-catalogue" className="list-group-item list-group-item-action">Manage Catalogue</Link>
                  <Link to="/graphql" className="list-group-item list-group-item-action">GraphQL</Link>
                </>
              }  
            </div>
          </nav>
        </div>
        
        <div className="col-md-10">
          <Alert message={alertMessage} className={alertClassName}/>
          <Outlet context={ {jwtToken, setJwtToken, setAlertMessage, setAlertClassName} }/>
        </div>
      </div>
      
    </div>
  );
}

export default App;
