import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';


// routes
import { app_routing } from './app.routes';
// Services

// Components TODO : move all declarations to their onws modules (is necesary create it)
import { AppComponent } from './app.component';
import { EditorComponent } from '../pages/editor/editor.component';
import { NavBar } from '../commons/components/nav-bar/nav-bar.component';
import { TextEditor } from '../commons/components/text-editor/text-editor.component';
import { LineEditor } from '../commons/components/text-editor/line/line.component';

@NgModule({
  declarations: [
    AppComponent,
    EditorComponent,
    NavBar,
    TextEditor,
    LineEditor
  ],
  imports: [
    BrowserModule,
    app_routing
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
