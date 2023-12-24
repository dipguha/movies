import Ticket from './../images/movie_tickets.jpg'

const Home = () => {

    return (
        <>
            <div className="text-center">
                <h2>Home screen - We are going to watch a Movie tonight</h2>
                <hr />
                <img src={Ticket} alt="Movie Tickets"></img>
            </div>
        </>
    )
}

export default Home;
