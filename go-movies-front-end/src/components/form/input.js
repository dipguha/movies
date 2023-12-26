import { forwardRef } from "react";

//forwardRef is a function from React library. Allows to forward ref to a child component or DOM element created inside a functional component
const Input = forwardRef( (props, ref) => {
    return (
        <div className="mb-3">
            {/* htmlFor is used to connect the label with the input field */}
            <label htmlFor={props.name} className="form-label">
                {props.title}
            </label>

            <input
                type={props.type}
                className={props.className}
                id={props.id}
                ref={ref}
                name={props.name}
                placeholder={props.placeholder}
                onChange={props.onChange}
                autoComplete={props.autoComplete}
                value={props.value}
            />

            <div className="">{props.errorMsg}</div>
        </div>
    )
})

export default Input;
