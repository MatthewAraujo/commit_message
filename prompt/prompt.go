package prompt

const response = `
:sparkles: feat: Integra OpenAI para geração de commits

Integradas bibliotecas para utilizar o OpenAI e gerar automaticamente
mensagens de commit. Alterações incluem:

- Inclusão dos pacotes openai e prompt.
- Função 'createUserMessage' para formatação do prompt.
- Uso do cliente OpenAI para gerar mensagens.
- Ajuste de variáveis 'task' e implementação do contexto.
`

func GetResponse() string {
	return response
}

const prompt = `
Commit Message Standards 📜
According to the Conventional Commits documentation, semantic commits are a simple convention used for commit messages. This convention defines a set of rules for creating an explicit commit history, facilitating the development of automated tools.

These commits help you and your team easily understand what changes were made to the committed code.

This identification occurs through a word and an emoji that specify whether the commit involves code changes, package updates, documentation, UI changes, testing, etc.

Type and Description 🦄
Semantic commits have the structural elements below (types), which indicate the intent of your commit to the user of your code.

feat - Commits of type feat indicate that your code introduces a new feature (related to MINOR in semantic versioning).

fix - Commits of type fix indicate that your committed code resolves a problem (bug fix) (related to PATCH in semantic versioning).

docs - Commits of type docs indicate changes to documentation, such as updates to your repository's README. (Does not include code changes.)

test - Commits of type test are used for changes to tests, such as creating, modifying, or removing unit tests. (Does not include code changes.)

build - Commits of type build are used for modifications to build files and dependencies.

perf - Commits of type perf are used to identify code changes related to performance improvements.

style - Commits of type style indicate changes related to code formatting, semicolons, trailing spaces, lint fixes, etc. (Does not include code changes.)

refactor - Commits of type refactor refer to changes due to refactoring that do not alter functionality, such as changes in how a certain screen is processed but maintaining the same functionality, or performance improvements from a code review.

chore - Commits of type chore indicate updates to build tasks, admin configurations, packages, etc., such as adding a package to .gitignore. (Does not include code changes.)

ci - Commits of type ci indicate changes related to continuous integration.

raw - Commits of type raw indicate changes related to configuration files, data, features, or parameters.

cleanup - Commits of type cleanup are used to remove commented code, unnecessary snippets, or any other source code cleanup, improving readability and maintainability.

remove - Commits of type remove indicate the deletion of obsolete or unused files, directories, or features, reducing project size and complexity while keeping it organized.

Recommendations 🎉
- Add a consistent type to the title.
- Limit the first line to a maximum of 4 words.
- Use the commit description to provide details.
- Use an emoji at the beginning of the commit message to represent its context.
- Add links in their original form—no link shorteners or affiliate links.

Commit Complements 💻
- Footer: information about the reviewer and card number in Trello or Jira. Example: Reviewed-by: Elisandro Mello Refs #133
- Body: precise descriptions of what the commit includes, impacts, and the reasons for the changes, as well as essential instructions for future interventions. Example: see the issue for details on typos fixed.
- Descriptions: a succinct description of the change. Example: correct minor typos in code.

💻 Examples
Git Command Result on GitHub
git commit -m "Initial commit" 🎉 Initial commit
git commit -m "docs: Update README" 📚 docs: Update README
git commit -m "fix: Infinite loop on line 50" 🐛 fix: Infinite loop on line 50
git commit -m "feat: Login page" ✨ feat: Login page
git commit -m "ci: Dockerfile modification" 🧱 ci: Dockerfile modification
git commit -m "refactor: Refactor to arrow functions" ♻️ refactor: Refactor to arrow functions
git commit -m "perf: Improved response time" ⚡ perf: Improved response time
git commit -m "fix: Revert inefficient changes" 💥 fix: Revert inefficient changes
git commit -m "feat: CSS styling for form" 💄 feat: CSS styling for form
git commit -m "test: Create new test" 🧪 test: Create new test
git commit -m "docs: Comments on LoremIpsum( ) function" 💡 docs: Comments on LoremIpsum( ) function
git commit -m "raw: RAW data for year YYYY" 🗃️ raw: RAW data for year YYYY
git commit -m "cleanup: Remove commented code and unused variables in form validation function" 🧹 cleanup: Remove commented code and unused variables in form validation function
git commit -m "remove: Remove unused project files for better organization and maintenance" 🗑️ remove: Remove unused project files for better organization and maintenance
`

func GetPrompt() string {
	return prompt
}
