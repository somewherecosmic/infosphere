import { FormEvent } from "react";
import { extractDataFromFormEvent } from "../util/DataHandling";

const Signup = () => {

    function onSubmitHandler(event: FormEvent<HTMLFormElement>) {
        const data = extractDataFromFormEvent(event);
    }

    function onResetHandler(event: FormEvent<HTMLFormElement>) {
        event.currentTarget.reset();
    }

    return (
        <>
            <form onSubmit={onSubmitHandler}>
                <label htmlFor="email">Email:</label>
                <input type="email"/>

                <label htmlFor="username">Username:</label>
                <input type="username" />

                <label htmlFor="password">Password:</label>
                <input type="password" />

                <button type="button">Reset</button>
                <button type="submit">Create Account</button>
            </form>
        </>
    );
}

export default Signup;