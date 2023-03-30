
import { Router, RouterModule, Routes } from '@angular/router';
import { HomePageComponent } from './home-page/home-page.component';
import { MainComponent } from './profile-page/main/main.component';

const appRoutes: Routes = [
  {path: 'home', component: HomePageComponent},
  {path: 'profile', component: MainComponent},
  {path: '', pathMatch: 'full', redirectTo: 'home'}
]

// @NgModule({
//   imports: [RouterModule.forRoot(routes)],
//   exports: [RouterModule]
// })
export const routing = RouterModule.forRoot(appRoutes);
