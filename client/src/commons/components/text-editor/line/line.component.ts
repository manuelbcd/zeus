import { Component, OnInit, Input, Output, EventEmitter, SimpleChanges} from '@angular/core';

import { LineClass } from '../../../classes/document/line.class';
@Component({
  selector: 'line-editor',
  templateUrl: './line.component.html',
  styleUrls: ['./line.component.css']
})
export class LineEditor {
    @Input() line: LineClass;

    // @Output() pageChange: EventEmitter<number> = new EventEmitter<number>();
    constructor() {

    }


}
