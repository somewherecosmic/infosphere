import React, { FormEvent } from "react";
import { extractDataFromFormEvent } from "../util/DataHandling";
import { useMutation } from "@tanstack/react-query";
import { publicAxios } from '../api/axiosInstances';

const Signup = () => {

    const mutation = useMutation({
        mutationFn: (formData: {[k:string]: FormDataEntryValue }) => {
            return publicAxios.post("/signup", formData, {
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded"
                }
            });
        }
    });

    function onSubmitHandler(event: FormEvent<HTMLFormElement>) {
        const formData = extractDataFromFormEvent(event);
        mutation.mutate(formData);
    }

    function onResetHandler(event: React.MouseEvent<HTMLButtonElement, MouseEvent>) {
        const target = event.currentTarget as HTMLElement;
        const form = target.parentElement as HTMLFormElement;
        form.reset();
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

                <button type="button" onClick={onResetHandler}>Reset</button>
                <button type="submit">Create Account</button>
            </form>

            {mutation.isPending && <p>Sending request...</p>}

            {mutation.isError && (<p>Error occurred: {mutation.error.message}</p>)}
            {mutation.isSuccess && <p>Account created</p>}
        </>
    );
}

export default Signup;