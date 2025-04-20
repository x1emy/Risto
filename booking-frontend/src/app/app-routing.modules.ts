import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RegisterComponent } from './register/register.component';
import { HomeComponent } from './home/home.component';
import {RestarauntsComponent} from './restaraunts/restaraunts.component';
import {RefugioComponent} from './refugio/refugio.component';
import {HaragunComponent} from './haragun/haragun.component';
import {AizaComponent} from './aiza/aiza.component';
import {BestiaComponent} from './bestia/bestia.component';

const routes: Routes = [
  { path: '', redirectTo: 'register', pathMatch: 'full' },
  { path: 'register', component: RegisterComponent },
  { path: 'home', component: HomeComponent },
  {path: 'restaraunts', component: RestarauntsComponent},
  { path: 'refugio', component:RefugioComponent },
  { path: 'haragun', component: HaragunComponent},
  { path: 'aiza', component: AizaComponent},
  { path: 'bestia', component: BestiaComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {}

