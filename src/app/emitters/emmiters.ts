import { EventEmitter } from "@angular/core";
import { Movie } from "../user-auth/user";

export class Emmiters {
    static authEmmiter = new EventEmitter<boolean>()
    static userData = new EventEmitter<string>()
    static isLoginOpen = new EventEmitter<boolean>()
    static watchList = new EventEmitter<Movie[]>()
}