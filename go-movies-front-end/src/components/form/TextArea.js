const TextArea = (props) => {
    return (
        
        <div className="mb-3">
            <label htmlFor={props.name} className="form-label">
                {props.title}
            </label>

            <textarea
                className="form-control"
                id={props.name}
                name={props.name}
                value={props.value}
                rows={props.rows}
                onChange={props.onChange}
            />
            <div className={props.errorDiv}>{props.errorMsg}</div>
        </div>
    )
}

export default TextArea;
