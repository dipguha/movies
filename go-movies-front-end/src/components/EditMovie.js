import { useEffect, useState } from "react";
import TextArea from "./form/TextArea";
import { useNavigate, useOutletContext, useParams } from "react-router-dom";
import Input from "./form/input";
import Select from "./form/Select";

const EditMovie = () => {

    const navigate = useNavigate();
    const { jwtToken } = useOutletContext();
    const [error, setError] = useState(null);
    const [errors, setErrors] = useState([]);

    // get id from the url
    let {id} = useParams();

    const mpaaOptions = [
        {id: "G", value: "G"},
        {id: "PG", value: "PG"},
        {id: "PG13", value: "PG13"},
        {id: "R", value: "R"},
        {id: "NC17", value: "NC17"},
        {id: "18A", value: "18A"},
    ]

    useEffect( () => {
        if (jwtToken === "") {
            navigate("/")
            return;
        }
    },[jwtToken, navigate])

    const hasError = (key) => {
        return errors.indexOf(key) !== -1;
    }

    const [movie, setMovie] = useState({
        id: 0,
        title: "",
        release_date: "",
        runtime: "",
        mpaa_rating: "",
        description: "",
    })

    const handleSubmit = (event) => {
        event.preventDefault();
    }

    // curried function
    const handleChange = () => (event) => {
        let value = event.target.value;
        let name = event.target.name;
        setMovie({
            ...movie,
            [name]: value,
        })
    }

    return (
        <div>
            <h2>EditMovie screen</h2>
            <pre>{JSON.stringify(movie, null, 3)}</pre>
            <hr />

            <form onSubmit={handleSubmit}>
                <input type="hidden" name="id" id="id" value={movie.id}></input>

                <Input
                    title={"Title"}
                    type={"text"}
                    name={"title"}
                    value={movie.title}
                    className={"form-control"}
                    onChange={handleChange("title")}
                    errorMsg={"Please enter a title"}
                    errorDiv={hasError("description") ? "text-danger" : "d-none"}                                        
                />

                <Input
                    title={"Release Data"}
                    type={"date"}
                    name={"release_date"}
                    value={movie.release_date}
                    className={"form-control"}
                    onChange={handleChange("release_date")}
                    errorMsg={"Please enter a release data"}
                    errorDiv={hasError("release_date") ? "text-danger" : "d-none"}                                        
                />

                <Input
                    title={"Runtime"}
                    className={"form-control"}
                    type={"text"}
                    name={"runtime"}
                    value={movie.runtime}
                    onChange={handleChange("runtime")}
                    errorDiv={hasError("runtime") ? "text-danger" : "d-none"}
                    errorMsg={"Please enter a runtime"}
                />

                <Select
                    title={"MPAA Rating"}
                    name={"mpaa_rating"}
                    options={mpaaOptions}
                    onChange={handleChange("mpaa_rating")}
                    placeholder={"Choose Rating"}
                    errorMsg={"Please enter a mpaa rating"}
                    errorDiv={hasError("mpaa_rating") ? "text-danger" : "d-none"}                                        
                />

                <TextArea 
                    title={"Description"}
                    name={"description"}
                    value={movie.description}
                    rows={"2"}
                    onChange={handleChange("description")}
                    errorMsg={"Please enter a description"}
                    errorDiv={hasError("description") ? "text-danger" : "d-none"}
                />


            </form>
        </div>
    )
}

export default EditMovie;
