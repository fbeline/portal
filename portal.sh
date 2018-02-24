function p() {
	  local output="$(portal $1)"
    if [[ -d "${output}" ]]; then
        cd "${output}"
    else
        echo "portal: directory '${@}' not found"
        echo "Try \`portal --help\` for more information."
        false
    fi
}

PROMPT_COMMAND='portal s $(pwd)'

