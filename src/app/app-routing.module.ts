import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomePageComponent } from './home-page/home-page.component';
import { ProfilePageModule } from './profile-page/profile-page.module';
import { MainComponent } from './profile-page/main/main.component';

const appRoutes: Routes = [
  { path: 'home', component: HomePageComponent },
  { path: 'profile', component: MainComponent },
  { path: '', pathMatch: 'full', redirectTo: 'home' }
]

// @NgModule({
//   imports: [RouterModule.forRoot(appRoutes)],
//   exports: [RouterModule]
// })
// export class AppRoutingModule { }

export const routing = RouterModule.forRoot(appRoutes);
