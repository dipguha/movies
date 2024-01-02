import { useCallback, useEffect, useState } from "react";
import { Link, Outlet, useNavigate } from "react-router-dom";
import Alert from "./components/Alert";

function App() {
  
  const [jwtToken, setJwtToken] = useState("");
  const [alertMessage, setAlertMessage] = useState("");
  const [alertClassName, setAlertClassName] = useState("d-none");
  const [tickInterval, setTickInterval] = useState();
  
  const navigate = useNavigate();

  //console.log("***** App.js jwtToken: ", jwtToken)

  //====================================================================
  const logout = () => {
    console.log("****** App.js-logout-jwtToken: ",jwtToken)

    const requestOptions = {
      method: "GET",
      credentials: "include"
    }

    fetch("/logout", requestOptions)
      .catch( error => {
        console.log("error logout: ", error)
      })
      .finally(() => {
        setJwtToken("")
        toggleRefresh(false);
      })
    
    navigate("/login")
  }

//====================================================================
const toggleRefresh = useCallback((status) => {
  console.log("***** toggleRefresh clicked: ");
  //600000
  if (status) {
    console.log("***** toggleRefresh turning on ticking status: ", status);
    let i  = setInterval(() => {

      const requestOptions = {
        method: "GET",
        credentials: "include",
      }

      fetch(`/refresh`, requestOptions)
      .then((response) => response.json())
      .then((data) => {
        if (data.access_token) {
          setJwtToken(data.access_token);
        }
      })
      .catch(error => {
        console.log("***** toggleRefresh-user is not logged in");
      })
    }, 600000);
    setTickInterval(i);
    console.log("setting tick interval to", i);
  } else {
    console.log("turning off ticking");
    console.log("turning off tickInterval", tickInterval);
    setTickInterval(null);
    clearInterval(tickInterval);
  }
}, [tickInterval])




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

      //console.log("****** App.js-useEffect-calling /refresh-requestOptions: ", requestOptions)
  
      fetch(`/refresh`, requestOptions)
      .then( (response) => {
        //console.log("response: ", response)
        //console.log("response status: ", response.status)
        //console.log("response headers: ", response.headers)
        //console.log("response body: ", response.body)
        //console.log("response Cookie: ", response.headers.get('Cookie'))

        if (response.status === 204) {
          console.log("Empty response received: ")  
        }
        
        return response.json()
      })
      .then( (data) => {
        //console.log("***** App.js-refresh-data: ", data)
        if (data.access_token) {
          setJwtToken(data.access_token);
          toggleRefresh(true);
        }
      })
      .catch(error => {
        console.log("****** App.js-useEffect-user is not logged in /refresh: ", error);
      })
    }
    console.log("****** App.js-useEffect-after /refresh data: ")
  }, [jwtToken, toggleRefresh])

  //console.log("***** App.js-just outside return: ")

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
          <Outlet context={ {jwtToken, setJwtToken, setAlertMessage, setAlertClassName, toggleRefresh,} }/>
        </div>
      </div>
      
    </div>
  );
}

export default App;
