import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

const Movies = () => {

    //useState is a hook in functional component to manage state
    const [movies, setMovies] = useState([])

    //useEffect is a hook in functional component to side effects (e.g. fetch data). Uses 3 lifecycle methods(DidMount, DidUpdate, WillUnmount)
    useEffect( () => {
        const headers = new Headers();
        headers.append("Content-Type", "application/json")

        const requestOptions = {
            method: "GET",
            headers: headers,
        }

        fetch(`http://localhost:8080/movies`, requestOptions)
            .then((response) => response.json())
            .then((data) => {
                setMovies(data)
            })
            .catch(err => {
                console.log(err)
            })
    }, [] )
    
    console.log("***** Movies.js *****")

    return (
        <div>
            <h2>Movies</h2>
            <hr />
            {/*Javascript expressions and templates used below along with map function */}
            <table className="table table-stripped table-hover">
                <thead>
                    <tr>
                        <th>Movies</th>
                        <th>Release Date</th>
                        <th>Rating</th>
                    </tr>
                </thead>
                <tbody>
                    {
                        movies.map( (m) => (
                            <tr key={m.id}>
                                <td>
                                    <Link to={`/movies/${m.id}`}>{m.title}</Link>
                                </td>
                                <td>{m.release_date}</td>
                                <td>{m.mpaa_rating}</td>
                            </tr>
                        ))
                    }
                </tbody>
            </table>
        </div>
    )
}

export default Movies;
