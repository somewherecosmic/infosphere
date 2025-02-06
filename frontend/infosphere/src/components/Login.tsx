import { FormEvent, FormEventHandler } from "react";
import { extractDataFromFormEvent } from "../util/DataHandling";

const Login = () => {

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
                <input name="email" type="email" />

                <label htmlFor="password">Password:</label>
                <input name="password" type="password" />

                <button type='button' onClick={onResetHandler}>Reset</button>
                <button type='submit'>Log in</button>
            </form>
        </>
    );
}

export default Login;