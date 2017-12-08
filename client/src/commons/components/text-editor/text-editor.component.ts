import { Component, OnInit } from '@angular/core';

import { DocumentClass } from '../../classes/document/document.class';

@Component({
  selector: 'text-editor',
  templateUrl: './text-editor.component.html',
  styleUrls: ['./text-editor.component.css']
})
export class TextEditor {
    document: DocumentClass;
    constructor() {
        this.document = new DocumentClass();
    }
}
