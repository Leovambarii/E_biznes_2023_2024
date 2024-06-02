import axios from "axios";

const Api = axios.create({
    headers: {
        "Content-Type": "application/json",
    },
    baseURL: "https://cloudservershopapp.azurewebsites.net",
});

export default Api;
