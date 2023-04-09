import { EventEmitter } from "@angular/core";

export class Emmiters {
    static authEmmiter = new EventEmitter<boolean>()
    static userData = new EventEmitter<string>()
    static isLoginOpen = new EventEmitter<boolean>()
}