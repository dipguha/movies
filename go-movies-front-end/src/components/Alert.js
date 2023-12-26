const Alert = (props) => {
    console.log("***** Alert.js *****")
    return(
        <div className={"alert " + props.className} role="alert">
            {props.message}
        </div>
    )
}

export default Alert;
