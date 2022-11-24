import axios from "axios";
import { env } from '$env/dynamic/public';

export const axiosInstance = axios.create({
    baseURL: env.PUBLIC_BASE_URL,
});