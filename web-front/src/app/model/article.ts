import { User } from './user';
export class Article {
    constructor(
        public id = "",
        public title = "",
        public body = "",
        public filename = "",
        public filedata = "",
        public created_datetime = "",
    ){ }
}