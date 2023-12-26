import { Link } from 'react-router-dom';
import Ticket from './../images/movie_tickets.jpg'

const Home = () => {

    return (
        <>
            <div className="text-center">
                <h2>Home screen - We are going to watch a Movie tonight</h2>
                <hr />
                <Link to="/movies">
                    <img src={Ticket} alt="Movie Tickets"></img>
                </Link>
            </div>
        </>
    )
}

export default Home;
