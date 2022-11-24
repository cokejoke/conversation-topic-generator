import type {Topic} from "./topic";

export interface Pack {
    id?: number,
    name: string,
    author: string,
    imageUrl?: string
    language: string,
    topics?: Array<Topic>
}