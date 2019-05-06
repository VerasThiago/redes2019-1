<!DOCTYPE html>

<html>
 

<body>
  <header>
    <h1 class="logo">XORA</h1>
    <div class="description">
      {{.firstName}}<br>
      {{.handle}}<br>
      {{.city}}<br>
      {{.organization}}<br>
      {{.maxRank}}<br>
    </div>
  </header>
  <br>
  <br>
  <form id="user">
    handle：<input name="username" type="text" />
    <input type="submit" value="submit" />
  </form>

  <script src="/static/js/reload.min.js"></script>
</body>
</html>
