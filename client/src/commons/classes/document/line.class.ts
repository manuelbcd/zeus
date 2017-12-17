import { Injectable } from '@angular/core';

import { LetterClass } from './letter.class';
@Injectable()
export class LineClass {
    //la linea llevara un array de letras, estado de la linea y el numero
    // la linea se encargara de parsetar el texto a un array
    line :number;
    text :string;
    constructor(line? :number, text?: string) {
        this.line = line;
        this.text = text;
    }
}