  // /wp-content/plugins/wp-postratings/js/postratings-js.js
  var post_id = 0,
      post_rating = 0,
      is_being_rated = !1;
  ratingsL10n.custom = parseInt(ratingsL10n.custom);
  ratingsL10n.max = parseInt(ratingsL10n.max);
  ratingsL10n.show_loading = parseInt(ratingsL10n.show_loading);
  ratingsL10n.show_fading = parseInt(ratingsL10n.show_fading);

  function current_rating(a, b, c) {
      if (!is_being_rated) {
          post_id = a;
          post_rating = b;
          if (ratingsL10n.custom && 2 == ratingsL10n.max) jQuery("#rating_" + post_id + "_" + b).attr("src", eval("ratings_" + b + "_mouseover_image.src"));
          else
              for (i = 1; i <= b; i++) ratingsL10n.custom ? jQuery("#rating_" + post_id + "_" + i).attr("src", eval("ratings_" + i + "_mouseover_image.src")) : jQuery("#rating_" + post_id + "_" + i).attr("src", ratings_mouseover_image.src);
          jQuery("#ratings_" + post_id + "_text").length && (jQuery("#ratings_" + post_id + "_text").show(), jQuery("#ratings_" +
              post_id + "_text").html(c))
      }
  }

  function ratings_off(a, b, c) {
      if (!is_being_rated) {
          for (i = 1; i <= ratingsL10n.max; i++) i <= a ? ratingsL10n.custom ? jQuery("#rating_" + post_id + "_" + i).attr("src", ratingsL10n.plugin_url + "/images/" + ratingsL10n.image + "/rating_" + i + "_on." + ratingsL10n.image_ext) : jQuery("#rating_" + post_id + "_" + i).attr("src", ratingsL10n.plugin_url + "/images/" + ratingsL10n.image + "/rating_on." + ratingsL10n.image_ext) : i == b ? ratingsL10n.custom ? jQuery("#rating_" + post_id + "_" + i).attr("src", ratingsL10n.plugin_url + "/images/" + ratingsL10n.image + "/rating_" +
              i + "_half" + (c ? "-rtl" : "") + "." + ratingsL10n.image_ext) : jQuery("#rating_" + post_id + "_" + i).attr("src", ratingsL10n.plugin_url + "/images/" + ratingsL10n.image + "/rating_half" + (c ? "-rtl" : "") + "." + ratingsL10n.image_ext) : ratingsL10n.custom ? jQuery("#rating_" + post_id + "_" + i).attr("src", ratingsL10n.plugin_url + "/images/" + ratingsL10n.image + "/rating_" + i + "_off." + ratingsL10n.image_ext) : jQuery("#rating_" + post_id + "_" + i).attr("src", ratingsL10n.plugin_url + "/images/" + ratingsL10n.image + "/rating_off." + ratingsL10n.image_ext);
          jQuery("#ratings_" +
              post_id + "_text").length && (jQuery("#ratings_" + post_id + "_text").hide(), jQuery("#ratings_" + post_id + "_text").empty())
      }
  }

  function set_is_being_rated(a) {
      is_being_rated = a
  }

  function rate_post_success(a) {
      jQuery("#post-ratings-" + post_id).html(a);
      ratingsL10n.show_loading && jQuery("#post-ratings-" + post_id + "-loading").hide();
      ratingsL10n.show_fading ? jQuery("#post-ratings-" + post_id).fadeTo("def", 1, function() {
          set_is_being_rated(!1)
      }) : set_is_being_rated(!1)
  }

  function rate_post() {
      post_ratings_el = jQuery("#post-ratings-" + post_id);
      if (is_being_rated) alert(ratingsL10n.text_wait);
      else {
          post_ratings_nonce = jQuery(post_ratings_el).data("nonce");
          if ("undefined" == typeof post_ratings_nonce || null == post_ratings_nonce) post_ratings_nonce = jQuery(post_ratings_el).attr("data-nonce");
          set_is_being_rated(!0);
          ratingsL10n.show_fading ? jQuery(post_ratings_el).fadeTo("def", 0, function() {
              ratingsL10n.show_loading && jQuery("#post-ratings-" + post_id + "-loading").show();
              jQuery.ajax({
                  type: "POST",
                  xhrFields: {
                      withCredentials: !0
                  },
                  dataType: "html",
                  url: ratingsL10n.ajax_url,
                  data: "action=postratings&pid=" + post_id + "&rate=" + post_rating + "&postratings_" + post_id + "_nonce=" + post_ratings_nonce,
                  cache: !1,
                  success: rate_post_success
              })
          }) : (ratingsL10n.show_loading && jQuery("#post-ratings-" + post_id + "-loading").show(), jQuery.ajax({
              type: "POST",
              xhrFields: {
                  withCredentials: !0
              },
              dataType: "html",
              url: ratingsL10n.ajax_url,
              data: "action=postratings&pid=" + post_id + "&rate=" + post_rating + "&postratings_" + post_id + "_nonce=" + post_ratings_nonce,
              cache: !1,
              success: rate_post_success
          }))
      }
  };
