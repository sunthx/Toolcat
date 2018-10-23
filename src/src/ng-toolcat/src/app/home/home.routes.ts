import { GenerateGuidComponent } from '../generate-guid/generate-guid.component';
import { HomeComponent } from '../home/home.component';
import { GenerateQrComponent } from '../generate-qr/generate-qr.component';
import { TimeConverterComponent } from '../time-converter/time-converter.component';
import { EncryptionComponent } from '../encryption/encryption.component';

export const HomeRoutes = [
    {
        path: '',
        component: HomeComponent,
        children: [
            {
                path: 'qr',
                component: GenerateQrComponent
            },
            {
                path: 'time',
                component: TimeConverterComponent
            },
            {
                path: 'encryption',
                component: EncryptionComponent
            },
            {
                path: 'uuid',
                component: GenerateGuidComponent
            },
            {
                path: '**',
                component: GenerateGuidComponent
            }
        ]
    },
    {
        path: '**',
        component: HomeComponent
    }
];
