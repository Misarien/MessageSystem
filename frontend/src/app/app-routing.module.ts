import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { ChatComponent } from './chat/chat.component';
import { SignupComponent } from './signup/signup.component';
import { HomeComponent } from './home/home.component';
const routes: Routes = [
  { path: '',
    component:HomeComponent,

  },
  {
    path: "login",component:LoginComponent
  },
  {
    path: "chat",
    component:ChatComponent,
    // loadChildren:'./chat'
  },
  {
    path: "signup",component:SignupComponent
  },
 // {path:"home",component:HomeComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
