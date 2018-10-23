import { HomeComponent } from './home/home.component';
import { AboutComponent } from './about/about.component';

export const AppRoutes = [
    {
        path: '',
        redirectTo: 'home',
        pathMatch: 'full'
    },
    {
        path: 'home',
        loadChildren: './home/home.module#HomeModule'
    },
    {
        path: 'about',
        component: AboutComponent
    },
    {
        path: '**',
        redirectTo: 'home',
        pathMatch: 'full'
    }
];
