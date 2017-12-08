import { RouterModule, Routes } from '@angular/router';
// import { HomeComponent } from './pages/home/home.component';
import { EditorComponent } from '../pages/editor/editor.component';

export const app_routes: Routes = [
    // { path: 'home', component: HomeComponent },
    { path: 'home', component: EditorComponent },
    // { path: 'detail/:id', component: DetailComponent },
    { path: '**', pathMatch: 'full', redirectTo: 'home' }
];

export const app_routing = RouterModule.forRoot(app_routes, { useHash: true });