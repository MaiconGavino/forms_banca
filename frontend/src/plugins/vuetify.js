import { createVuetify } from "vuetify";
import "vuetify/styles";

export default createVuetify({
    theme: {
        themes: {
            light: {
                colors: {
                    primary: "#3f51b5",
                    secondary: "#4caf50",
                    accent: "#ff9800",
                    background: "#f5f5f5",
                    surface: "#ffffff",
                    text: "#212121",
                },
            },
        },
    },
});
