import { Routes } from '@angular/router';
import { RegisterComponent } from './register/register.component';
import { HomeComponent } from './home/home.component';
import {RestarauntsComponent} from './restaraunts/restaraunts.component';

export const appRoutes: Routes = [
  { path: '', redirectTo: 'register', pathMatch: 'full' },
  { path: 'register', component: RegisterComponent },
  { path: 'home', component: HomeComponent },
  { path: 'restaraunts', component: RestarauntsComponent}
];
