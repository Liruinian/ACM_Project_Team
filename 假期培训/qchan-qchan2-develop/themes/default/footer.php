<?php

/* Deny direct visit */
if (!defined('INDEX_RUN')) {
	header('HTTP/1.1 403 Forbidden');
	exit('This file must be loaded in flow.');
}

?>

<!-- Footer -->
<footer id="main_footer">
	<p><a href="https://beian.miit.gov.cn/" target="_blank">黑ICP备2023001389号</a></p>
	<?= format_main_site() ?>
</footer>
</body>

</html>