import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

const Movies = () => {

    //useState is a hook in functional component to manage state
    const [movies, setMovies] = useState([])

    //useEffect is a hook in functional component to side effects (e.g. fetch data). Uses 3 lifecycle methods(DidMount, DidUpdate, WillUnmount)
    useEffect( () => {
        let movieList = [
            { id: 1, title: "Movie 1", release_date: "01-01-2001", runtime: 111, mpaa_rating: "R1", description: "Description 1" },
            { id: 2, title: "Movie 2", release_date: "02-02-2002", runtime: 222, mpaa_rating: "R2", description: "Description 2" },
            { id: 3, title: "Movie 3", release_date: "03-03-2003", runtime: 333, mpaa_rating: "R3", description: "Description 3" },
        ];
        setMovies(movieList)
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
