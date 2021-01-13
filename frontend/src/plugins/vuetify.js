import Vue from "vue";
import Vuetify from "vuetify/lib/framework";

Vue.use(Vuetify);

export default new Vuetify({
  theme: {
    themes: {
      light:  {
        primary: '#D5222C',
        secondary:'#13151E',
        accent: '#D5C3BE',
        error: '#f80a18',
        warning: '#E0B724',
        info: '#3AA797',
        success: '#5B8B29'
        },
    },
  },
})