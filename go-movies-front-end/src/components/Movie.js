import { useEffect, useState } from "react";
import { useParams } from "react-router";

const Movie = () => {

    //use empty object as an initial value
    const[movie, setMovie] = useState({});
    //useParams comes from react-router
    let { id } = useParams();

    useEffect( () => {
        let myMovie = { id: 1, title: "Movie 1", release_date: "01-01-2001", runtime: 111, mpaa_rating: "R1", description: "Description 1" }
        setMovie(myMovie)
    }, [id] )

    return (
        <div>
            <h2>Movie: {movie.title}</h2>
            <small><em>{movie.release_date}, {movie.runtime} minutes, Rated {movie.mpaa_rating}</em></small>
            <hr />
            <p>{movie.description}</p>
        </div>
    )
}

export default Movie;
