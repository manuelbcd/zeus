import { Injectable } from '@angular/core';

@Injectable()
export class LetterClass {
    /*letter class, esta es chunga joder, que llevara la muy puta, pues lleva una letra, estado o situacion
        El evento tambien A
        */
    letter: string;
    constructor(letter: string) {
        this.letter = letter;
    }
}
