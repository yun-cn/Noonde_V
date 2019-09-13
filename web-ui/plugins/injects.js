import { CancelToken } from 'axios';
import moment          from 'moment-timezone'
import Cookie          from 'js-cookie'

export default ({ app, store }, inject) => {
  inject('wsURL', app.$env.WS_URL || 'ws://127.0.0.1:3000');

  inject('source', () => {
    return CancelToken.source()
  });

  inject('pre', (val) => {
    if (!val) return
    return val.split('\n').join('<br/>')
  });

  inject('map', (arr, key) => {
    return [].concat(arr || []).map(x => x[key])
  });

  inject('uniq', (arr) => {
    return [].concat(arr || [])
      .filter((v, i, a) => a.indexOf(v) == i)
  });

  inject('include', (arr, val) => {
    return [].concat(arr || [])
      .filter((v, i, a) => a.indexOf(v) == i)
      .map(x => String(x))
      .includes(String(val))
  });

  inject('actives', (arr) => {
    return [].concat(arr || [])
      .filter((v, i, a) => a.indexOf(v) == i)
      .filter(x => x.active)
      .map(x => x.value)
  });

  inject('remove', (arr, val) => {
    return [].concat(arr || [])
      .filter((v, i, a) => a.indexOf(v) == i)
      .filter(x => String(x) != String(val))
  });

  inject('page', (from, per) => {
    return Math.floor(from / per) + 1
  });

  inject('length', (total, per) => {
    if (!total || total == 0) return 1;
    return Math.floor((total - 1) / per) + 1
  });

  inject('format', (time, fmt) => {
    let t = moment(time);
    return t.format(fmt);
  });

  inject('color', store.state.color);

  inject('style', () => {
    let style = {};
    for (let key in store.state.color) {
      let dashed = '--' + key.replace(/[A-Z]/g, m => '-' + m.toLowerCase());
      style[dashed] = store.state.color[key];
    }

    return style
  });

  inject('auth', {
    getFromCookie (req) {
      if (!req.headers.cookie) return;

      let cookies = req.headers.cookie.split(';');
      if (!cookies) return;

      let token            = cookies.find(c => c.trim().startsWith('token='         ));
      let id               = cookies.find(c => c.trim().startsWith('id='            ));
      let email            = cookies.find(c => c.trim().startsWith('email='         ));
      let nickname         = cookies.find(c => c.trim().startsWith('nickname='      ));
      let avatar           = cookies.find(c => c.trim().startsWith('avatar='        ));
      let profile          = cookies.find(c => c.trim().startsWith('profile='       ));
      // let urlName          = cookies.find(c => c.trim().startsWith('urlName='       ));
      let noteCount        = cookies.find(c => c.trim().startsWith('noteCount='     ));
      let followerCount    = cookies.find(c => c.trim().startsWith('followerCount=' ));
      let followingCount   = cookies.find(c => c.trim().startsWith('followingCount='));

      let data = {};
      data.id             = id                 && id.split('=')[1];
      data.token          = token              && token.split('=')[1];
      data.email          = email              && email.split('=')[1];
      data.nickname       = nickname           && nickname.split('=')[1];
      data.avatar         = avatar             && avatar.split('=')[1];
      data.profile        = profile            && profile.split('=')[1];
      // data.urlName        = urlName            && urlName.split('=')[1];
      data.noteCount      = noteCount          && noteCount.split('=')[1];
      data.followerCount  = followerCount      && followerCount.split('=')[1];
      data.followingCount = followingCount     && followingCount.split('=')[1];
      return data
    },

    getFromLocalStorage () {
      let data             = {};
      data.token           = window.localStorage.token;
      data.id              = window.localStorage.id;
      data.email           = window.localStorage.email;
      data.nickname        = window.localStorage.nickname;
      data.avatar          = window.localStorage.avatar;
      data.profile         = window.localStorage.profile;
      // data.urlName         = window.localStorage.urlName;
      data.noteCount       = window.localStorage.noteCount;
      data.followerCount   = window.localStorage.followerCount;
      data.followingCount  = window.localStorage.followingCount;
      return data
    },

    set (token, user, remember) {
      if (process.SERVER_BUILD) return;
      console.log('Set');

      window.localStorage.setItem('token',           token               );
      window.localStorage.setItem('id',              user.id             );
      window.localStorage.setItem('email',           user.email          );
      window.localStorage.setItem('nickname',        user.nickname       );
      window.localStorage.setItem('avatar',          user.avatar         );
      window.localStorage.setItem('profile',         user.profile        );
      // window.localStorage.setItem('urlName',         user.url_name       );
      window.localStorage.setItem('noteCount',       user.note_count     );
      window.localStorage.setItem('followerCount',   user.follower_count );
      window.localStorage.setItem('followingCount',  user.following_count );

      Cookie.set('token',            token,                { expires: remember ? 3650 : null });
      Cookie.set('id',               user.id,              { expires: remember ? 3650 : null });
      Cookie.set('email',            user.email,           { expires: remember ? 3650 : null });
      Cookie.set('nickname',         user.nickname,        { expires: remember ? 3650 : null });
      Cookie.set('avatar',           user.avatar,          { expires: remember ? 3650 : null });
      Cookie.set('profile',          user.profile,         { expires: remember ? 3650 : null });
      // Cookie.set('urlName',          user.url_name,        { expires: remember ? 3650 : null });
      Cookie.set('noteCount',        user.note_count,      { expires: remember ? 3650 : null });
      Cookie.set('followerCount',    user.follower_count,  { expires: remember ? 3650 : null });
      Cookie.set('followingCount',   user.following_count, { expires: remember ? 3650 : null });
    },

    unset () {
      if (process.SERVER_BUILD) return;

      window.localStorage.removeItem('token' );
      window.localStorage.removeItem('id'    );
      window.localStorage.removeItem('email' );
      window.localStorage.removeItem('name'  );
      window.localStorage.removeItem('avatar');
      window.localStorage.removeItem('profile');
      // window.localStorage.removeItem('urlName');
      window.localStorage.removeItem('noteCount');
      window.localStorage.removeItem('followerCount');
      window.localStorage.removeItem('followingCount');

      Cookie.remove('token' );
      Cookie.remove('id'    );
      Cookie.remove('email' );
      Cookie.remove('name'  );
      Cookie.remove('avatar');
      Cookie.remove('profile');
      // Cookie.remove('urlName');
      Cookie.remove('noteCount');
      Cookie.remove('followerCount');
      Cookie.remove('followingCount');
    },
  })
}
