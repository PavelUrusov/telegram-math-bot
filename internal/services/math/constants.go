package math

const DefaultTimeoutUpdate = 60
const DefaultCacheDuration = 30
const DefaultOffsetUpdate = 0

// commands endpoints
const (
	CalculateCommandID     = "calculate"
	PlotCommandID          = "plot"
	SolveEquationCommandID = "solve"
	StartCommandID         = "start"
	HelpCommandID          = "help"
)
const startCommandResponse = `
Привет! 👋 Я ваш математический помощник-бот. Вот что я могу для вас сделать:

1. /calculate <выражение> - Выполняю элементарные математические вычисления. Просто отправьте мне выражение, и я быстро предоставлю ответ. Например, попробуйте /calculate 2+2.

2. /plot <функция> - Строю графики функций. Отправьте мне функцию, и я создам для вас её график. Например, используйте /plot sin(x) для построения графика функции синуса.

3. /solve <уравнение> - Решаю уравнения. Дайте мне уравнение, и я найду его корни. Например, команда /solve x^2-4=0 поможет вам найти решение этого квадратного уравнения.

4. /help - Получить информацию о том, как использовать бота и описание всех доступных команд.

5. /start - Начать работу с ботом и увидеть это приветственное сообщение.

Чтобы начать, просто введите одну из этих команд с вашим запросом. Если у вас возникнут вопросы или вам потребуется помощь, просто напишите мне в любое время!
`

const invalidCommandResponse = `
Извините, я не распознаю эту команду. 😕 Пожалуйста, используйте одну из следующих доступных команд:

1. /calculate <выражение> - Для выполнения математических вычислений. Например, /calculate 3+4*2.

2. /plot <функция> - Чтобы получить график заданной функции. Например, /plot x^2.

3. /solve <уравнение> - Для решения математических уравнений. Например, /solve x^2-9=0.

4. /help - Вызвать это справочное сообщение.

5. /start - Показать приветственное сообщение и основную информацию о боте.

Убедитесь, что вы вводите команду корректно. Если вам нужна дополнительная помощь или у вас есть вопросы, не стесняйтесь обращаться!
`

const errorResponse = `
Извините, я не смог обработать ваш запрос. 🤔 Возможные причины:

1. Введенные данные некорректны или не полны. Пожалуйста, проверьте введенные данные и попробуйте снова.

2. Запрос слишком сложен или не поддается вычислению моими средствами.

Если вы уверены, что данные верны, и нуждаетесь в дополнительной помощи, попробуйте уточнить ваш запрос или обратитесь за помощью, используя команду /help.
`
