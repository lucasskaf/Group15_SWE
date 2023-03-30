import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ProfilePageModule } from './profile-page/profile-page.module';

const routes: Routes = [
  // {path: 'profile', loadChildren: () => ProfilePageModule}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
