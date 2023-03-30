import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomePageComponent } from './home-page/home-page.component';
import { ProfilePageModule } from './profile-page/profile-page.module';

const routes: Routes = [
  {path: '', component: HomePageComponent},
  {path: 'profile', component: ProfilePageModule}
]

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
