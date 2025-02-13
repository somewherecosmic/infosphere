import axios from 'axios'


export const publicAxios = axios.create({
    baseURL: "http://localhost:8080/api"
});


// const privateAxios = axios.create({
//     baseURL: 
// })