import { Injectable } from '@angular/core';

import { LineClass } from './line.class';

@Injectable()
export class DocumentClass {
    //El document se va a componer de lineas, usuario, fecha de modificación ,....
    //array de lineas

    lines: any[];
    constructor() {
        this.loadMock();
    }

    loadMock() {
        this.lines = [
            new LineClass(1, 'texto en linea 1'),
            new LineClass(2, 'texto en linea 2'),
            new LineClass(3, 'texto en linea 3'),
            new LineClass(4, 'texto en linea 4'),
            new LineClass(5, 'texto en linea 5'),
            new LineClass(6, 'texto en linea 6'),
            new LineClass(7, 'texto en linea 7'),
            new LineClass(8, 'texto en linea 8'),
            new LineClass(9, 'texto en linea 9'),
            new LineClass(10, 'texto en linea 10'),
            new LineClass(11, 'texto en linea 11'),
            new LineClass(12, 'texto en linea 12'),
            new LineClass(13, 'texto en linea 13'),
            new LineClass(15, 'texto en linea 14'),
            new LineClass(16,'texto en linea 15'),
            new LineClass(17, 'texto en linea 16'),
            new LineClass(18, 'texto en linea 17')];

        }
}
